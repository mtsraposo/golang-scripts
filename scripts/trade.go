package scripts

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	. "github.com/mtsraposo/golang-scripts/utils"
)

type TradeDataFrame struct {
	dataframe.DataFrame
}

type TradeDataFrameGroups struct {
	*dataframe.Groups
}

// Trade Using int for monetary calculations to avoid rounding errors
// Trade A scale of precision (e.g. 1e6) may be used for conversion.
type Trade struct {
	TradeDate  string
	Instrument string
	SideIsBuy  bool
	Quantity   int
	Price      int
}

type Total struct {
	TradeDate string
	Total     int
}

type Position struct {
	Instrument string
	Total      int
}

func TradeSummary(trades []Trade) TradeDataFrame {
	df := &TradeDataFrame{dataframe.LoadStructs(trades)}
	return *df.appendTotals(true).
		groupBy([]string{"TradeDate"}).
		agg([]string{"TradeDate"}).
		castToInt().
		sortBy("TradeDate")
}

func TradePositions(trades []Trade, date string) TradeDataFrame {
	df := &TradeDataFrame{dataframe.LoadStructs(trades)}
	return *df.appendTotals(false).
		groupBy([]string{"Instrument"}).
		accumulatePositions().
		filterOutEmptyPositions().
		filterByDate(date).
		sortBy("Instrument").
		selectCols([]string{"Instrument", "Total"})
}

func (df *TradeDataFrame) appendTotals(cash bool) *TradeDataFrame {
	colNames := df.Names()
	_totals := func(row series.Series) series.Series {
		return totalsByRow(row, colNames, cash)
	}
	totals := df.Rapply(_totals).Rename("Total", "X0")
	return &TradeDataFrame{
		df.CBind(totals),
	}
}

func totalsByRow(row series.Series, colNames []string, cash bool) series.Series {
	total := calcTotal(row, colNames, cash)
	return series.Ints(total)
}

func calcTotal(row series.Series, colNames []string, cash bool) int {
	sideMultiplier := extractSideMultiplier(colNames, row, cash)
	quantityIndex := IndexOf("Quantity", colNames)
	quantity, haveExtractedQuantity := row.Elem(quantityIndex).Int()
	Must(haveExtractedQuantity)
	if cash {
		priceIndex := IndexOf("Price", colNames)
		price, haveExtractedPrice := row.Elem(priceIndex).Int()
		Must(haveExtractedPrice)
		return price * quantity * sideMultiplier
	}
	return quantity * sideMultiplier
}

func extractSideMultiplier(colNames []string, row series.Series, cash bool) int {
	sideIndex := IndexOf("SideIsBuy", colNames)
	sideIsBuy, haveExtractedBool := row.Elem(sideIndex).Bool()
	Must(haveExtractedBool)
	if (sideIsBuy && cash) || (!sideIsBuy && !cash) {
		return -1
	}
	return 1
}

func (df *TradeDataFrame) groupBy(cols []string) *TradeDataFrameGroups {
	return &TradeDataFrameGroups{df.GroupBy(cols...)}
}

func (df *TradeDataFrameGroups) agg(cols []string) *TradeDataFrame {
	return &TradeDataFrame{df.
		Aggregation([]dataframe.AggregationType{dataframe.Aggregation_SUM},
			[]string{"Total"}).
		Rename("Total", "Total_SUM").
		Select(append(cols, []string{"Total"}...)),
	}
}

func (df *TradeDataFrame) castToInt() *TradeDataFrame {
	colNames := df.Names()
	colsToKeep := filterOutColsToKeep(colNames, "Total")
	_castToInt := func(row series.Series) series.Series {
		return castRowToInt(row, colNames)
	}
	totals := df.Rapply(_castToInt).Rename("Total", "X0")
	return &TradeDataFrame{
		df.Select(colsToKeep).CBind(totals),
	}
}

func filterOutColsToKeep(colNames []string, colToFilter string) []string {
	var colsToKeep []string
	for _, col := range colNames {
		if col != colToFilter {
			colsToKeep = append(colsToKeep, col)
		}
	}
	return colsToKeep
}

func castRowToInt(row series.Series, colNames []string) series.Series {
	totalIndex := IndexOf("Total", colNames)
	total := row.Elem(totalIndex).Float()
	return series.Ints(total)
}

func (df *TradeDataFrame) sortBy(col string) *TradeDataFrame {
	return &TradeDataFrame{
		df.Arrange(
			dataframe.Sort(col),
		),
	}
}

func (df *TradeDataFrameGroups) accumulatePositions() *TradeDataFrame {
	cols := []string{"TradeDate", "Instrument"}
	var cumulativeTotals, cumulativeInstrument, aggByDate, totalsSeries dataframe.DataFrame
	for _, aggDf := range df.GetGroups() {
		aggByDate = groupByAndAggregate(aggDf, []string{"TradeDate", "Instrument"}).
			Arrange(dataframe.Sort("TradeDate"))
		totalsSeries = aggByDate.
			Select("Total").
			Capply(accumulateTotals)
		cumulativeInstrument = aggByDate.Select(cols).CBind(totalsSeries)
		cumulativeTotals = cumulativeTotals.Concat(cumulativeInstrument)
	}
	return &TradeDataFrame{cumulativeTotals}
}

func groupByAndAggregate(df dataframe.DataFrame, cols []string) dataframe.DataFrame {
	return df.
		GroupBy(cols...).
		Aggregation([]dataframe.AggregationType{dataframe.Aggregation_SUM},
			[]string{"Total"}).
		Rename("Total", "Total_SUM").
		Select(append(cols, []string{"Total"}...))
}

func accumulateTotals(col series.Series) series.Series {
	currentTotal, haveExtractedTotal := col.Elem(0).Int()
	Must(haveExtractedTotal)
	rows := col.Len()
	cumulativeTotals := []int{currentTotal}
	for i := 1; i < rows; i++ {
		currentTotal, haveExtractedTotal = col.Elem(i).Int()
		Must(haveExtractedTotal)
		cumulativeTotals = append(cumulativeTotals, cumulativeTotals[i-1]+currentTotal)
	}
	return series.New(cumulativeTotals, series.Int, "Total")
}

func (df *TradeDataFrame) filterOutEmptyPositions() *TradeDataFrame {
	return &TradeDataFrame{
		df.Filter(
			dataframe.F{
				Colname:    "Total",
				Comparator: series.Neq,
				Comparando: 0,
			},
		),
	}
}

func (df *TradeDataFrame) filterByDate(date string) *TradeDataFrame {
	return &TradeDataFrame{
		df.Filter(
			dataframe.F{
				Colname:    "TradeDate",
				Comparator: series.Eq,
				Comparando: date,
			},
		),
	}
}

func (df *TradeDataFrame) selectCols(cols []string) *TradeDataFrame {
	return &TradeDataFrame{
		df.Select(cols),
	}
}

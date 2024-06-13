import ChartDataQueryBuilder from "./ChartDataQueryBuilder";
import ChartDataFetcher from "$lib/data-fetcher/ChartDataFetcher";
import DoughnutChartData from "$lib/model/DoughnutChartData";
import LineChartData from "$lib/model/LineChartData";
import BubbleChartData from "$lib/model/BubbleChartData";
import DataParameters from "$lib/model/DataParameters";
import ParameterDataFetcher from "$lib/data-fetcher/ParameterDataFetcher";

// Service for the dashboard controller.
class DashboardService {
  // Holds the dependency injected query builder instance.
  #queryBuilder;
  // Holds the dependency injected chart data fetcher.
  #chartDataFetcher;
  // Holds the dependency injected parameter data fetcher.
  #parameterDataFetcher;
  constructor(
    queryBuilder = new ChartDataQueryBuilder(),
    chartDataFetcher = new ChartDataFetcher(),
    parameterDataFetcher = new ParameterDataFetcher(),
  ) {
    this.#queryBuilder = queryBuilder;
    this.#chartDataFetcher = chartDataFetcher;
    this.#parameterDataFetcher = parameterDataFetcher;
  }

  // Queries the backend and construct model entity with result
  // based on the type of query.
  async getChartData(formData) {
    const queryString = this.#createQueryString(formData);
    const data = await this.#chartDataFetcher.getChartData(queryString);
    if (formData.aggregationChoice == "total") {
      const doughtnutChartData = new DoughnutChartData(data);
      return doughtnutChartData;
    }
    if (formData.aggregationChoice == "by year") {
      const lineChartData = new LineChartData(data);
      return lineChartData;
    }
    if (formData.aggregationChoice == "all games") {
      const bubbleChartData = new BubbleChartData(data);
      return bubbleChartData;
    }
  }

  // Private helper method for constructing an appropiate 
  // query string based on the user selections.
  #createQueryString(formData) {
    const qb = this.#queryBuilder.new();
    qb.aggregate(formData.aggregationChoice);
    qb.grouping(formData.groupingChoice);
    if (formData.years) {
      qb.startYear(formData.years.startYear);
      qb.endYear(formData.years.endYear);
    }
    if (formData.selectedYear) {
      qb.startYear(formData.selectedYear);
      qb.endYear(formData.selectedYear);
    }
    if (formData.aggregationChoice == 'by year') {
      qb.filters(formData.filtering);
    }
    return qb.build();
  }

  // Method for getting all valid parameters.
  async getValidParameters() {
    const result = await this.#parameterDataFetcher.getParameters();
    const parameters = new DataParameters(result);
    return parameters;
  }
}

export default DashboardService;

// Represents a chart data query builder, used to create valid
// query params that can be used to query the endpoint with. Implements
// the builder pattern.
class ChartDataQueryBuilder {
  // Variable holding the url params.
  #params
  constructor() {
    this.#params = new URLSearchParams();
  }

  // Method for adding filter query param.
  // Example: "filters=filter1,filter2,filter3".
  filters(values) {
    const filtersStr = values.join(",")
    this.#addParam("filters", filtersStr);
  }

  // Method for adding end-year query param.
  endYear(value) {
    this.#addParam("end-year", value);
  }

  //Method for adding start-year query param.
  startYear(value) {
    this.#addParam("start-year", value);
  }

  // Method adding grouping query param.
  grouping(value) {
    this.#addParam("group", value);
  }

  // Method for adding aggregate query param.
  aggregate(value) {
    this.#addParam("aggregate", value);
  }

  // Method for adding a search/query query param.
  search(value) {
    this.#addParam("query", value)
  }

  // Method for adding page query param.
  page(value) {
    this.#addParam("page", value)
  }

  // Private method for creating a query param and adding it to the 
  // params variable.
  #addParam(key, value) {
    let lowerCaseVal;
    if (typeof value == 'number') {
      lowerCaseVal = value.toString();
    } else {
      lowerCaseVal = value.toLowerCase();
    }
    this.#params.append(key, lowerCaseVal);
  }

  // Builds the query param string.
  build() {
    return this.#params.toString();
  }

  // Returns a new chart data query builder. 
  new() {
    return new ChartDataQueryBuilder();
  }
}

export default ChartDataQueryBuilder;



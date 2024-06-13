import SearchDataFetcher from "$lib/data-fetcher/SearchDataFetcher";
import ChartDataQueryBuilder from "./ChartDataQueryBuilder";
import SearchResult from "$lib/model/SearchResult";

// Service for the search controller.
class SearchService {
  // Holds the dependency injected query builder.
  #queryBuilder;
  // Holds the dependency injected data fetcher.
  #dataFetcher;

  constructor(dataFetcher = new SearchDataFetcher(), queryBuilder = new ChartDataQueryBuilder()) {
    this.#dataFetcher = dataFetcher;
    this.#queryBuilder = queryBuilder;
  }

  // Queries the backend and returns a model.SearchResult object.
  async searchData(searchString, page) {
    if (searchString != "") {
      const queryStr = this.#createQueryString(searchString, page);
      const data = await this.#dataFetcher.getSearchData(queryStr);
      const dataResponse = new SearchResult(data);
      return dataResponse;
    }
  }

  // Private helper method for creating a query string.
  #createQueryString(searchString, page) {
    const qb = this.#queryBuilder.new();
    qb.search(searchString);
    qb.page(page);
    return qb.build();
  }
}

export default SearchService;

import { PUBLIC_SEARCHDATA_ENDPOINT, PUBLIC_API_KEY } from '$env/static/public';

// Represents a data fetcher that queries the search endpoint for
// search query data.
class SearchDataFetcher {
  // The endpoint for querying the search data.
  #endpointURL;
  // The api key to be provided in the request.
  #apiKey;

  constructor() {
    this.#endpointURL = PUBLIC_SEARCHDATA_ENDPOINT;
    this.#apiKey = PUBLIC_API_KEY;
  }

  // Method for getting the data from the endpoint.
  async getSearchData(queryString) {
    try {
      const url = `${this.#endpointURL}?${queryString}`;
      console.log("Fetching data from: " + url);
      const response = await fetch(url, {
        method: "GET",
        headers: {
          "X-API-Key": this.#apiKey,
          "Content-Type": "application/json",
        }
      });

      if (!response.ok) {
        throw new Error(`Error fetching data: ${response.statusText}`)
      }

      const data = await response.json();
      return data;
    } catch (error) {
      console.error("failed to fetch search data: ", error);
      throw error;
    }
  }
}

export default SearchDataFetcher;

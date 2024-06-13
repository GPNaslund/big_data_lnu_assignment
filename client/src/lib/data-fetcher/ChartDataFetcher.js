import { PUBLIC_CHARTDATA_ENDPOINT, PUBLIC_API_KEY } from '$env/static/public';

// Represents a data fetcher for the aggregated chart data to display.
class ChartDataFetcher {
  // The endpoint to query for chart data.
  #endpointURL;

  // The api key to be provided in the request.
  #apiKey;

  constructor() {
    this.#endpointURL = PUBLIC_CHARTDATA_ENDPOINT;
    this.#apiKey = PUBLIC_API_KEY;
  }

  // Method for querying the endpoint for chart data.
  async getChartData(queryString) {
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
      console.error("failed to fetch chart data: ", error);
      throw error;
    }
  }
}

export default ChartDataFetcher;

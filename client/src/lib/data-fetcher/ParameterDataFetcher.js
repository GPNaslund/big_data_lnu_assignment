import { PUBLIC_PARAMETERS_ENDPOINT, PUBLIC_API_KEY } from '$env/static/public';

// Class with methods for fetching the data parameters that can be used
// to query the backend for data.
class ParameterDataFetcher {
  // Holds the endpoint url.
  #endpoint;
  // Holds the api key.
  #apiKey;

  // Constructs a new instance of ParameterDataFetcher.
  constructor() {
    this.#endpoint = PUBLIC_PARAMETERS_ENDPOINT;
    this.#apiKey = PUBLIC_API_KEY;
  }

  // Method for querying the endpoint for the valid data parameters.
  async getParameters() {
    try {
      const url = `${this.#endpoint}`;
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

export default ParameterDataFetcher;

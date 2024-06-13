import VideoGame from "./VideoGame";

// Model entity containing the data from search query response.
class SearchResult {
  // Variable for holding the amount of documents found.
  #found;
  // An array of video games entities.
  #videoGamesData;

  constructor(dataResponse) {
    this.#videoGamesData = [];
    this.addVideoGamesData(dataResponse)
    this.#found = dataResponse.found;
  }

  // Method for populating the video games data variable with video games entities.
  addVideoGamesData(dataResponse) {
    dataResponse.Hits.forEach((hit) => {
      this.#videoGamesData.push(new VideoGame({
        rank: hit.document.Rank,
        name: hit.document.Name,
        platform: hit.document.Platform,
        year: hit.document.Year,
        genre: hit.document.Genre,
        publisher: hit.document.Publisher,
        naSales: hit.document.NaSales,
        euSales: hit.document.EuSales,
        jpSales: hit.document.JpSales,
        otherSales: hit.document.OtherSales,
        globalSales: hit.document.GlobalSales,
      }));
    });
  }

  // Getter for the array of video games data.
  get videoGamesData() {
    return this.#videoGamesData;
  }

  // Getter for the found variable.
  get found() {
    return this.#found;
  }
}

export default SearchResult;

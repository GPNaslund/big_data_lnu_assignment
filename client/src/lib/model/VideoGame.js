// Model entity representing a video game. Contains various info 
// about the video game.
class VideoGame {
  // The video games ranking of total sales in the dataset.
  #rank;
  // The video game name
  #name;
  // The platform the video game was released on.
  #platform;
  // The year the video game was released.
  #year;
  // The genre of the video game.
  #genre;
  // The publisher of the video game.
  #publisher;
  // The total amount of sales in North America.
  #naSales;
  // The total amount of sales in Europe.
  #euSales;
  // The total amount of sales in Japan.
  #jpSales;
  // The total amount of sales in other regions.
  #otherSales;
  // The total global amount of sales.
  #globalSales;

  constructor({ rank, name, platform, year, genre, publisher, naSales, euSales, jpSales, otherSales, globalSales }) {
    this.#rank = rank;
    this.#name = name;
    this.#platform = platform;
    this.#year = year;
    this.#genre = genre;
    this.#publisher = publisher;
    this.#naSales = naSales;
    this.#euSales = euSales;
    this.#jpSales = jpSales;
    this.#otherSales = otherSales;
    this.#globalSales = globalSales;
  }

  // Getter for the rank.
  get rank() {
    return this.#rank;
  }

  // Getter for the name.
  get name() {
    return this.#name;
  }

  // Getter for the platform.
  get platform() {
    return this.#platform;
  }

  // Getter for the year.
  get year() {
    return this.#year;
  }

  // Getter for the genre.
  get genre() {
    return this.#genre;
  }

  // Getter for the publisher.
  get publisher() {
    return this.#publisher;
  }

  // Getter for the North America sales.
  get naSales() {
    return this.#naSales;
  }

  // Getter for the Europe sales.
  get euSales() {
    return this.#euSales;
  }

  // Getter for the Japan sales.
  get jpSales() {
    return this.#jpSales;
  }

  // Getter for the sales of other regions.
  get otherSales() {
    return this.#otherSales;
  }

  // Getter for the global sales.
  get globalSales() {
    return this.#globalSales;
  }
}

export default VideoGame;

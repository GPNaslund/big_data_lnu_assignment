// Represents all valid data parameters.
class DataParameters {
  #aggregations;
  #groupings;
  #regions;
  #genres;
  #startYear;
  #endYear;

  constructor(data) {
    this.#dataValidator(data);
    this.#aggregations = data.aggregations;
    this.#groupings = data.groupings;
    this.#regions = data.regions;
    this.#genres = data.genres;
    this.#startYear = data["start-year"];
    this.#endYear = data["end-year"];
  }

  // Getter for aggregations array.
  get aggregations() {
    return this.#aggregations;
  }

  // Getter for groupings array.
  get groupings() {
    return this.#groupings;
  }

  // Getter for regions array.
  get regions() {
    return this.#regions;
  }

  // Getter for genres array.
  get genres() {
    return this.#genres;
  }

  // Getter for valid start year.
  get startYear() {
    return this.#startYear;
  }

  // Getter for end year.
  get endYear() {
    return this.#endYear;
  }

  // Validates the data.
  #dataValidator(data) {
    if (data === undefined) {
      throw new Error("No data provided")
    }
    this.#dataPropertyArrayValidator("aggregations", data.aggregations);
    this.#dataPropertyArrayValidator("groupings", data.groupings);
    this.#dataPropertyArrayValidator("regions", data.regions);
    this.#dataPropertyArrayValidator("genres", data.genres);
    if (data["start-year"] === undefined) {
      throw new Error("Start year is undefined");
    }
    if (data["end-year"] === undefined) {
      throw new Error("End year is undefined");
    }
  }

  // Validates existence and length of data property.
  #dataPropertyArrayValidator(name, property) {
    if (property === undefined || property.length === 0) {
      throw new Error(name + " is undefined or empty")
    }
  }

}

export default DataParameters;

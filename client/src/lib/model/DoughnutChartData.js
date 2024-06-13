import { ColorPalette } from "./ColorPalette";

// Model entity holding the data used in the doughnut chart. 
// Used for the total sales.
class DoughnutChartData {
  // Variable holding the color palette.
  palette;
  // Variable holding the data set labels.
  labels;
  // Variable holding the datasets to be displayed.
  datasets;

  constructor(dataArray) {
    if (dataArray.length === 0) {
      throw new Error("No data in provided array to doughnut chart data");
    }

    // Sets the palette to the imported color palette.
    this.palette = ColorPalette;

    this.labels = dataArray.map(item => item.data.Category);
    this.datasets = [{
      data: dataArray.map(item => item.data.Sales),
      backgroundColor: this.generateColors(dataArray.length),
    }];
  }

  // Method for getting random colors from the color palette.
  generateColors(count) {
    let randomColors = [];
    for (let i = 0; i < count; i++) {
      const randomIndex = Math.floor(Math.random() * this.palette.length);
      randomColors.push(this.palette[randomIndex]);
    }
    return randomColors;
  }
}

export default DoughnutChartData;

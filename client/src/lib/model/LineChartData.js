import { ColorPalette } from "./ColorPalette";

// Model entity for holding the data to be displayed in the line chart.
// Used for the yearly data display.
class LineChartData {
  // Variable holding the color palette.
  palette;
  // Variable holding the dataset labels to be displayed.
  labels;
  // Variable containing the data points to be displayed.
  datasets;
  // Variable for holding the max sales for setting the y-axis roof.
  maxSales;

  constructor(dataArray) {
    if (!dataArray || dataArray.length === 0) {
      throw new Error("Data array is empty or undefined");
    }

    // Sets the color palette.
    this.palette = ColorPalette;
    this.labels = [];
    this.datasets = {};
    this.maxSales = 0;

    this.populateFields(dataArray);
  }

  // Method for populating the labels and datasets variables.
  populateFields(dataArray) {
    dataArray.forEach(entry => {
      const year = entry.data.Year;
      this.labels.push(year);

      entry.data.Data.forEach(item => {
        if (!this.datasets[item.Category]) {
          this.datasets[item.Category] = {
            label: item.Category,
            data: [],
            borderColor: this.getColor(),
            fill: false
          };
        }
        this.datasets[item.Category].data.push({
          x: year,
          y: item.Sales
        });

        if (item.Sales > this.maxSales) {
          this.maxSales = item.Sales;
        }
      });
    });

    this.normalizeData();
  }

  // Method for normalizing the dataset.
  normalizeData() {
    this.labels.sort((a, b) => a - b);
    this.datasets = Object.values(this.datasets).map(dataset => {
      const fullData = [];
      this.labels.forEach(year => {
        const dataPoint = dataset.data.find(point => point.x === year);
        fullData.push(dataPoint ? dataPoint.y : 0);
      });
      dataset.data = fullData;
      return dataset;
    });
  }

  // Method for getting a random color from the color palette.
  getColor() {
    const randomIndex = Math.floor(Math.random() * this.palette.length);
    return this.palette[randomIndex]
  }
}

export default LineChartData;

import { ColorPalette } from "./ColorPalette";

// Model entity for holding data to be displayed by bubble chart.
// Is used for displaying all released games by a specific year.
class BubbleChartData {
  // Variable for the color palette.
  palette;
  // Variable holding the max sales.
  maxSales;
  // The datasets to be displayed.
  datasets;

  constructor(dataArray) {
    if (dataArray.length === 0) {
      throw new Error("No data in provided array for bubble chart");
    }

    // Sets the color palette used for the chart dataset.
    this.palette = ColorPalette;
    // Variable for getting the global sales for each data item, used to normalize radius from
    // the largest bubble.
    const salesValues = dataArray.map(item => item.data.GlobalSales);
    // Variable for holding the largest sale number.
    const maxSales = Math.max(...salesValues);
    // Holds the minimum radius size for the bubbles.
    const minRadius = 1;
    // Holds the maxiumum radius size for the bubbles.
    const maxRadius = 10;

    this.datasets = [{
      label: 'Game Sales',
      data: dataArray.map(item => {
        const normalizedRadius = (Math.sqrt(item.data.GlobalSales / maxSales) * (maxRadius - minRadius)) + minRadius;
        const randomX = Math.random() * 100;
        return {
          x: randomX,
          y: item.data.GlobalSales,
          r: normalizedRadius,
          gameInfo: {
            Name: item.data.Name,
            Year: item.data.Year,
            Platform: item.data.Platform,
            NaSales: item.data.NaSales,
            EuSales: item.data.EuSales,
            JpSales: item.data.JpSales,
            OtherSales: item.data.OtherSales,
            GlobalSales: item.data.GlobalSales
          }
        };
      }),
      backgroundColor: this.generateColors(dataArray.length),
    }];
  }

  // Method for getting random colors from the color palette that is used for the datapoints in the chart.
  generateColors(count) {
    return Array.from({ length: count }, () => this.palette[Math.floor(Math.random() * this.palette.length)]);
  }
}

export default BubbleChartData;

# Wine quality

Abstract: Two datasets are included, related to red and white vinho verde wine samples, from the north of Portugal. The goal is to model wine quality based on physicochemical tests (see [Cortez et al., 2009]).
http://archive.ics.uci.edu/ml/datasets/Wine+Quality

## Data
|fixed acidity|volatile acidity|citric acid|residual sugar|chlorides|free sulfur dioxide|total sulfur dioxide|density|pH|sulphates|alcohol|__quality__|
|-------------|----------------|-----------|--------------|---------|------------------|-------------|------|--------|--|---------|-------|-------|
|7|0.27|0.36|20.7|0.045|45|170|1.001|3|0.45|8.8|6|
|6.3|0.3|0.34|1.6|0.049|14|132|0.994|3.3|0.49|9.5|6|
|8.1|0.28|0.4|6.9|0.05|30|97|0.9951|3.26|0.44|10.1|6|



## Example
This example demonstrate the application of a regressor:
* with 100 epochs, 70 percent training data, 0.9 learning with 0.001 decay
* Regression threshold of 0.2
* 100 hidden neurons

## Learning
The learning here is that the regressor decides between correct vs. wrong classified using a threshold. You can change this threshold to bring more tolerance to the system.

## Source
Paulo Cortez, University of Minho, Guimarães, Portugal, http://www3.dsi.uminho.pt/pcortez
A. Cerdeira, F. Almeida, T. Matos and J. Reis, Viticulture Commission of the Vinho Verde Region(CVRVV), Porto, Portugal
@2009

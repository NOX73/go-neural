# Sonar Data Set

Abstract: The task is to train a network to discriminate between sonar signals bounced off a metal cylinder and those bounced off a roughly cylindrical rock.

Connectionist Bench (Sonar, Mines vs. Rocks) Data Set

Found here: https://archive.ics.uci.edu/ml/datasets/Connectionist+Bench+%28Sonar%2C+Mines+vs.+Rocks%29

## Example

In this example a data set was used to demonstrate
* a MLP with 100 hidden neurons
* that uses CriterionDistance to decide for the best model
* gives a summary of the training
* and persists the file

Below the command line output can be seen.
```
> go run main.go

...

summary for class R
 * TP: 23 TN: 30 FP: 0 FN: 8
 * Recall/Sensitivity: 0.7419354838709677
 * Precision: 1
 * Fallout/FalsePosRate: 0
 * False Discovey Rate: 0
 * Negative Prediction Rate: 0.7894736842105263
--
 * Accuracy: 0.8688524590163934
 * F-Measure: 0.8518518518518519
 * Balanced Accuracy: 0.8709677419354839
 * Informedness: 0.7419354838709677
 * Markedness: 0.7894736842105263

summary for class M
 * TP: 30 TN: 23 FP: 8 FN: 0
 * Recall/Sensitivity: 1
 * Precision: 0.7894736842105263
 * Fallout/FalsePosRate: 0.25806451612903225
 * False Discovey Rate: 0.21052631578947367
 * Negative Prediction Rate: 1
--
 * Accuracy: 0.8688524590163934
 * F-Measure: 0.8823529411764706
 * Balanced Accuracy: 0.8709677419354839
 * Informedness: 0.7419354838709677
 * Markedness: 0.7894736842105263
```

// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Jakub Malhotra
// Created on: March 2023
// This file contains the JS functions for index.html

"use strict"
/**
 * This function calculates the area of a parallelogram
 */
function enterClicked() {
  //input
  const base = parseInt(document.getElementById("base-of-parallelogram").value)
  const height = parseInt(
    document.getElementById("height-of-parallelogram").value
  )

  //process
  const area = base * height

  //output
  document.getElementById("area-of-parallelogram").innerHTML =
    "The area of the parallelogram is " + area + " cm²."
}

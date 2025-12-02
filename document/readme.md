---
title: "Algorithms & Data-structures"
subtitle: "Hands-on implementation and analysis"
author: "Pieter Groenendijk (HAN 1654942)"
date: "Nov 18 2025"
lang: "en-US"
bibliography: "./references.bib"
geometry: "left=3cm,right=3cm,top=2cm,bottom=2cm"
header-includes:
    - \usepackage{attachfile}
---

\newpage
# Lists
## Dynamic array 
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

- In golang, the length of an array is part of it's type, meaning the length must be a constant. Only slices can
be created with a length specified at runtime, which already is a dynamic array...

- Not being able to specifically see if there is actually unallocated heap memory after the already defined underlying
array is something only the built-in slice implementation may know. For our implementation we don't have this granular
control and must reallocate a whole different array and copy the values upong growing.

\newpage
## Linked list
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
## Comparison
<!--
    - Je implementeert twee typen reeksen en maakt een vergelijking tussen de implementaties.

    - Je onderbouwt de verschillen tussen de twee implementaties en maakt hierbij gebruik 
    van de concepten time complexity en execution time. Je behandelt hierbij zowel het 
    slechtst mogelijke geval als het best mogelijke geval en legt uit waarin deze van elkaar verschillen.
-->

\newpage
# Priority Queue
<!--
    - Je implementeert een priority queue.
    - Je legt uit op welke manier je de priority implementeert en onderbouwt waarom je voor deze methode hebt gekozen.
    - Je onderbouwt de performance van je implementatie en maakt hierbij gebruik van de concepten 
    time complexity en execution time. Je behandelt hierbij zowel het slechtst mogelijke geval als het 
    best mogelijke geval en legt uit waarin deze van elkaar verschillen.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Sorting
## Merge sort
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

## Insertion sort
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

## Comparison
<!--
    - Je implementeert een merge sort algoritme en 1 ander sorteeralgoritme naar 
    keuze (NIET parallel merge sort) en maakt een vergelijking tussen de implementaties.
    - Je onderbouwt de verschillen tussen de twee implementaties en maakt 
    hierbij gebruik van de concepten time complexity, execution time en space complexity. Je 
    behandelt hierbij zowel het slechtst mogelijke geval als het best mogelijke geval en 
    legt uit waarin deze van elkaar verschillen.
--> 

\newpage
# Hash Table
<!--
    - Je implementeert een hash table.
    - Je legt uit hoe je hashing-algoritme werkt en onderbouwt je keuze voor het specifieke algoritme.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Binary Search
<!--
    - Je implementeert een binary search.
    - Je onderbouwt de performance van je implementatie en maakt hierbij gebruik van de concepten time 
    complexity en execution time. Je behandelt hierbij zowel het slechtst mogelijke geval als het 
    best mogelijke geval en legt uit waarin deze van elkaar verschillen.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->

\newpage
# Graphs
## [Graph implementation]
<!--
    - Je implementeert Dijkstra's algoritme. Hiervoor gebruik je een door jou 
    zelf ontwikkelde implementatie van een graaf.
    - Je onderbouwt de performance van je implementatie van een graaf en maakt hierbij 
    gebruik van de concepten time complexity en execution time.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->
## Dijkstra's
<!--
    - Je implementeert Dijkstra's algoritme. Hiervoor gebruik je een door jou zelf 
    ontwikkelde implementatie van een graaf.
    - Je onderbouwt de performance van je implementatie van Dijkstra's 
    algoritme en maakt hierbij gebruik van de concepten time complexity en execution time.
    - Je legt uit welke onderdelen van je code zich lenen voor verbeteringen die impact hebben op de performance.
-->


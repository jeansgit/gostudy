package main

import "fmt"

func main() {
 var a int = 21
 var b int = 10
 
 if(a==b){
 	fmt.Printf("a等于b\n")	
 }else{
 	fmt.Printf("a不等于b\n")
 }
 if(a<b){
 	fmt.Printf("a小于b\n")
 }else{
 	fmt.Printf("a不小于b\n")
 }
 if(a>b){
 	fmt.Printf("a大于b\n")
 }else {
 	fmt.Printf("a不大于b\n")
 }
 a = 5
 b = 20
 if(a>=b){
 	fmt.Printf("a大于等于b\n")
 }
 if(a<=b){
 	fmt.Printf("a小于等于b\n")
 }
 }
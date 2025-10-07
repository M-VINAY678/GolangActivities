## Marshaling And UnMarshaling
* First we need to take some dummy api, which produces dummy data 
* The api should also accept limit, so that we can get required amount of data.
* Then use net/http package for getting dummy data in json format
* Now I UnMarshaled it using two funtions
  * one function is used when you know about structure of the response or json data.
  * another function is used when you don't kown about structure of the response, we use empty interfaces.
* After unmarshalling data, we need to marshall data .
   * First created users.json file
   * then using marshal, converted []User datatype to json format and 
   * Then using os package stored in user.json file

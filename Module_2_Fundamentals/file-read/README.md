## Reading File using os packing
* First I took file path as user input.
* Then I used os package and read the file using os.ReadFile(Path)
* So if their is no error,Then I am directly printing to the terminal using fmt.Println(fileContent)

## Handling Error during File Read

* If their is Error
  * Then first I am checking whether it is related to file does not exist error using
  * errors.Is(err,os.ErrNotExist),In this case, i would print "the file does not exist".
  * Incase,if it is other than above error, then I am printing "Error during File Read"

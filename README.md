# Sungava
Static site generator build using [**go**](https://golang.org/).

### Prerequisites
You must have **go** installed in your local machine to get started.

* **golang**
	
  Installation process is documented [here](https://golang.org/doc/install).
  
* Set up the **GOPATH**

  It defaults to `$HOME/go`. But if you've initialized **go** in another directory then set the `GOPATH` as: 
  ```
  export GOPATH=$HOME/<path-to-your-directory>
  ```
  For more details on the `GOPATH`, run the following command on you terminal/shell:
  ```
  go help gopath
  ```
  
* For this project following dependencies for **go** must be installed
	* [gin-gonic/gin](https://github.com/gin-gonic/gin)
	* [gin-contrib.static](https://github.com/gin-contrib/static)
	* [russross/balckfirday](https://github.com/russross/blackfriday)

Make sure to have all dependencies install before getting started.

### Run
* Clone the repo inside go `$GOPATH/src/github.com/your-username/`

* Write the content of your webpage in markdown inside index.md which is located inside markdown folder and run:

  ```
  go run main.go
  ```

* After this go to `localhost:3000` you can see your webpage up and running.


*Note: Project is underdevelopment I want to build sungava something like a Jekyll*

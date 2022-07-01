function getMovies(){
    fetch("http://localhost:8080/movies")
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write("response:\n")
        document.write(str)
    });
}

function createMovie(){
    fetch("http://localhost:8080/movies", {
        method : 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body : JSON.stringify( {
            'isbn' : document.getElementById("createMovieIsbn").value,
            'title' : document.getElementById("createMovieTitle").value,
            'director' : {
                "firstname" : document.getElementById("createMovieName").value,
                "lastname" : document.getElementById("createMovieSurname").value
            }
        })
    })
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write("you created:\n")
        document.write(str)
    })
}
function updateMovie(id){
    fetch("http://localhost:8080/movies/" + id, {
        method : 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body : JSON.stringify( {
            'isbn' : document.getElementById("UpdateMovieIsbn").value,
            'title' : document.getElementById("UpdateMovieTitle").value,
            'director' : {
                "firstname" : document.getElementById("UpdateDirectorName").value,
                "lastname" : document.getElementById("UpdateDirectorSurname").value
            }
        })
    })
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write("you updated movie with id " + id + " to:\n")
        document.write(str)
    })
}

function getMovieById(id){
    fetch("http://localhost:8080/movies/" + id)
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write("response:\n")
        document.write(str)
    });
}

function deleteMovie(id){
    fetch("http://localhost:8080/movies/" + id, {
        method: 'DELETE',
    })
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write("you deleted:\n")
        document.write(str)
    })
}


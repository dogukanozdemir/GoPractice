function getMovies(){
    fetch("http://localhost:8080/movies")
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
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
            'name' : document.getElementById("createMovieTitle").value,
            'director' : {
                "firstname" : document.getElementById("createMovieName").value,
                "lastname" : document.getElementById("createMovieSurname").value
            }
        })
    })
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write(str)
    })
}
function updateMovie(id){
    fetch("http://localhost:8080/movies", {
        method : 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body : JSON.stringify( {
            'Id' : id,
            'isbn' : document.getElementById("UpdateMovieIsbn").value,
            'name' : document.getElementById("UpdateMovieTitle").value,
            'director' : {
                "firstname" : document.getElementById("UpdateDirectorName").value,
                "lastname" : document.getElementById("UpdateDirectorSurname").value
            }
        })
    })
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
        document.write(str)
    })
}

function getMovieById(id){
    fetch("http://localhost:8080/movies/" + id)
    .then(response => response.json())
    .then(data => {
        var str = JSON.stringify(data);
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
        document.write(str)
    })
}


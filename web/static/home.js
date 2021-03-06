// Ger the genres and fill in the dropdown
$.get( "/genres", function( data ) {
  for (i=0; i<data.length; i++) {
    optText = data[i];
    optValue = data[i];
    $('#genres').append(new Option(optText, optValue));
  }
});

// Get the list of current downloading files
$.get( "/current", function( data ) {

  if (data.length > 0) {
    $('.current-downloads').append("<tr><th>Status</th><th>Genre</th><th>URL</th><tr>")
  }

  for (i=0; i<data.length; i++) {
    status = data[i].Status
    genre = data[i].Genre
    id = data[i].Id

    if (id.includes("youtube")) {
        $('.current-downloads').append(`<tr><td>${status}</td><td>${genre}</td><td><a href='${id}'>${id}</a></td></tr>`);
    } else {
      var ytl = `youtube.com/watch?v=${id}`
    $('.current-downloads').append(`<tr><td>${status}</td><td>${genre}</td><td><a href='https://${ytl}'>${ytl}</a></td></tr>`);
        }

    }

    urlStr = "youtube.com/watch?v="

    $('.current-downloads').append(`<tr><td>${status}</td><td>${genre}</td><td>${urlStr}${id}</td></tr>`);
  }

});

// Get the full history of all downloaded files

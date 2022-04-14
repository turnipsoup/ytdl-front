$.get( "/genres", function( data ) {
  for (i=0; i<data.length; i++) {
    optText = data[i];
    optValue = data[i];
    $('#genres').append(new Option(optText, optValue));
  }
});
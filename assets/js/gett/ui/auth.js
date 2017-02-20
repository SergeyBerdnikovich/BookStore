$(document).ajaxError( function(event, xhr){
  if (xhr.status == 401) {
    console.log("Authentication failed, reloading the page...");
    window.location.reload();
  }
});


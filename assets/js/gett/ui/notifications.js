$(function() {
  "use strict";
  var bsAlert = function(alert_type, message) {

    var alertHTML = ['<div id="alertdiv" class="alert alert-',  alert_type,
      '"><a class="close" data-dismiss="alert">×</a><span>', message, '</span></div>'];
    $('#alert_placeholder').html(alertHTML.join(""));
  };

  var $body = $('body');
  $body.on('showAlert', function(e, alert_type, message){
    bsAlert(alert_type, message);
    setTimeout(function() {
      // automatically close the alert and remove this if the users doesnt close it
      $("#alertdiv").slideUp();
    }, 3000);
  });

  $body.on('showPermanentAlert', function(e, alert_type, message){
    bsAlert(alert_type, message);
  });
});

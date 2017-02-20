$.rails.allowAction = function(link) {
  if (!link.attr('data-confirm')) {
    return true;
  }
  $.rails.showConfirmDialog(link);
  return false;
};

$.rails.confirmed = function(link) {
  link.removeAttr('data-confirm');
  return link.trigger('click.rails');
};

$.rails.showConfirmDialog = function(link) {
  var modal = $('#confirmModal');
  modal.find('.modal-body h4').html(link.data('confirm'));
  modal.modal();
  return modal.find('#confirm').on('click', function() {
    modal.modal('hide');
    return $.rails.confirmed(link);
  });
};
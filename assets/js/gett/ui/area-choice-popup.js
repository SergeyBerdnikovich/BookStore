window.AreaChoicePopup = (function(){
  return function(root, onChange) {

    var areas = {};
    areas.popupOpened = false;
    areas.onChange = onChange || function(){};

    var showModal = function(content) {
      var newModal = $('#area-choice-popup').clone();
      newModal.find('.modal-body').html(content);
      areas.popupOpened = true;
      newModal.modal('show').on('hidden.bs.modal', function(){
        areas.popupOpened = false;
        newModal.remove();
      });
      return newModal;
    };

    areas.el = function() { return root.find('.table'); };
    areas.endpoint = function() { return 'http://' + areas.el().data('endpoint'); };

    var jwtKey = $('body').data('jwt-key');
    var jwt = $('body').data('jwt');

    if (jwt && jwtKey) {
      areas.signature = jwtKey + '=' + jwt;
    } else {
      areas.signature = '';
    }

    var env = location.pathname.match(/[^\/]+/);
    var search = areas.el().data('search') || '';
    areas.iframeHtml = function(){
      var ids = areas.ids(),
          url = areas.endpoint() + '/' + env + '/areas/selection?ids=' + ids.join(',') + '&search=' + encodeURIComponent(search) + '&' + areas.signature;

      return '<iframe src="' + url + '">'
    };

    areas.ids = function(){
      return areas.el()
        .find('tbody tr')
        .map(function(){ return $(this).data('id') })
        .toArray()
    };

    areas.getArea = function(id){
      return $.getJSON(areas.endpoint() + '/api/v1/' + env + '/areas/' + id);
    };

    areas.nameCache = {};

    areas.loadNameForElement = function($el){
      var id = $el.data('id');
      var $nameEl = $el.find('.name');
      if (!id || !$nameEl.length)
          return;

      if (areas.nameCache[id]) {
        $nameEl.text(areas.nameCache[id]);
      } else {
        areas.getArea(id).then(function(json){
          areas.nameCache[json.data.id] = json.data.title;
          $nameEl.text(json.data.title);
        });
      }
    };

    areas.loadNames = function(){
      areas.el().find('tbody tr').each(function(){
        areas.loadNameForElement($(this));
      });
    };

    areas.buffer = function(message){
      if (message.checked) {
        areas._buffer[message.id] = message;
        areas.nameCache[message.id] = message.name
      } else {
        delete areas._buffer[message.id];
      }
      areas._bufferChanged = true;
    };

    var addAreasSelector = '.add-area';
    var areaPrefix = function() { return areas.el().data('prefix'); };
    var areaParentId = function() { return areas.el().data('parentId'); };

    areas.resetBuffer = function(){
      this._buffer = {};
      this._bufferChanged = false;
    }

    areas.applyBuffer = function(){
      if (!areas._bufferChanged) return;

      var ids = Object.keys(areas._buffer).join(',');

      $.post((root.find(addAreasSelector).data('href') || 'areas'), { ids: ids, prefix: areaPrefix(), parent_id: areaParentId() }).then(function(response){
        var newData = $(response).find('tbody').html();
        areas.el().find('tbody').html(newData);
        areas.loadNames();
        areas.onChange(areas.ids());
      });

      this.resetBuffer();
    };

    root.on('click', addAreasSelector, function(){
      showModal(areas.iframeHtml()).on('hidden.bs.modal', function(){
        areas.applyBuffer();
      });
      return false;
    });

    $(window).on('message', function(e){
      if (areas.popupOpened) {
        areas.buffer(e.originalEvent.data);
      }
    });

    var deleteAreaSelector = 'a[data-method=delete]';
    // see gett-ui/app/assets/javascripts/gett/ui/confirmable.js
    root.on('click.rails', deleteAreaSelector, function() {
      var el = $(this)

      $.post((el.data('href') || el.attr('href')), { _method: 'delete', prefix: areaPrefix(), parent_id: areaParentId() }).then(function(){
        el.parents('tr').remove();
        areas.onChange(areas.ids());
      });
      return false;
    });

    areas.resetBuffer();
    return areas;
  };
})();

var app = angular.module('app', []);
 
function hostname(s) {
    var l = window.location;
    var path = l.pathname;
    //need to do better job here
    if(s[0] == "/")
      path = "";
    return ((l.protocol === "https:") ? "wss://" : "ws://") + l.hostname + (((l.port != 80) && (l.port != 443)) ? ":" + l.port : "") + path + s;
}

app.factory('ChatService', function() {
  var service = {};
 
  service.connect = function() {
    if(service.ws) { return; }
 
    var ws = new WebSocket(hostname("/entry"));
 
    ws.onopen = function() {
      service.callback('{"author":"CHANNEL","body":"Succeeded to open a connection"}');
    };
 
    ws.onerror = function() {
      service.callback('{"author":"CHANNEL","body":"Failed to open a connection"}');
    }
 
    ws.onmessage = function(message) {
      service.callback(message.data);
    };
 
    service.ws = ws;
  }
 
  service.send = function(message) {
    service.ws.send(message);
  }
 
  service.subscribe = function(callback) {
    service.callback = callback;
  }
 
  return service;
});
 
 
function AppCtrl($scope, ChatService) {
  $scope.messages = [];
 
  ChatService.subscribe(function(message) {
    try{
      var messageJSON=JSON.parse(message)
      $scope.messages.push(messageJSON);
      $scope.$apply();
    }catch(e){
      console.debug(e, message);
    }
  });
 
  $scope.connect = function() {
    ChatService.connect();
  }
 
  $scope.send = function() {
    ChatService.send('{"author": "abc", "body": "'+$scope.body+'"}');
    $scope.body = "";
  }
}

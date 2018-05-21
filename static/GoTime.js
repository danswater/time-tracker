// Generated by CoffeeScript 1.7.1
(function() {
  window.GoTime = (function() {
    GoTime._syncCount = 0;

    GoTime._offset = 0;

    GoTime._precision = Infinity;

    GoTime._history = [];

    GoTime._syncInitialTimeouts = [0, 3000, 9000, 18000, 45000];

    GoTime._syncInterval = 900000;

    GoTime._synchronizing = false;

    GoTime._lastSyncTime = null;

    GoTime._lastSyncMethod = null;

    GoTime._ajaxURL = null;

    GoTime._ajaxSampleSize = 1;

    GoTime._firstSyncCallbackRan = false;

    GoTime._firstSyncCallback = null;

    GoTime._onSyncCallback = null;

    GoTime._wsCall = null;

    GoTime._wsRequestTime = null;

    function GoTime() {
      GoTime._setupSync();
      return new Date(GoTime.now());
    }

    GoTime._setupSync = function() {
      var time, _i, _len, _ref;
      if (GoTime._synchronizing === false) {
        GoTime._synchronizing = true;
        _ref = GoTime._syncInitialTimeouts;
        for (_i = 0, _len = _ref.length; _i < _len; _i++) {
          time = _ref[_i];
          setTimeout(GoTime._sync, time);
        }
        setInterval(GoTime._sync, GoTime._syncInterval);
      }
    };

    GoTime.now = function() {
      return Date.now() + GoTime._offset;
    };

    GoTime.getOffset = function() {
      return GoTime._offset;
    };

    GoTime.getPrecision = function() {
      return GoTime._precision;
    };

    GoTime.getLastMethod = function() {
      return GoTime._lastSyncMethod;
    };

    GoTime.getSyncCount = function() {
      return GoTime._syncCount;
    };

    GoTime.getHistory = function() {
      return GoTime._history;
    };

    GoTime.setOptions = function(options) {
      if (options.AjaxURL != null) {
        GoTime._ajaxURL = options.AjaxURL;
      }
      if (options.SyncInitialTimeouts != null) {
        GoTime._syncInitialTimeouts = options.SyncInitialTimeouts;
      }
      if (options.SyncInterval != null) {
        GoTime._syncInterval = options.SyncInterval;
      }
      if (options.OnSync != null) {
        GoTime._onSyncCallback = options.OnSync;
      }
      if (options.WhenSynced != null) {
        GoTime._firstSyncCallback = options.WhenSynced;
      }
      return GoTime._setupSync();
    };

    GoTime.wsSend = function(callback) {
      return GoTime._wsCall = callback;
    };

    GoTime.wsReceived = function(serverTimeString) {
      var responseTime, sample, serverTime;
      responseTime = Date.now();
      serverTime = GoTime._dateFromService(serverTimeString);
      sample = GoTime._calculateOffset(GoTime._wsRequestTime, responseTime, serverTime);
      return GoTime._reviseOffset(sample, "websocket");
    };

    GoTime._ajaxSample = function() {
      var req, requestTime;
      req = new XMLHttpRequest();
      req.open("GET", GoTime._ajaxURL);
      req.onreadystatechange = function() {
        var responseTime, sample, serverTime;
        responseTime = Date.now();
        if (req.readyState === 4) {
          if (req.status === 200) {
            serverTime = GoTime._dateFromService(req.responseText);
            sample = GoTime._calculateOffset(requestTime, responseTime, serverTime);
            GoTime._reviseOffset(sample, "ajax");
          }
        }
      };
      requestTime = Date.now();
      req.send();
      return true;
    };

    GoTime._sync = function() {
      var success;
      if (GoTime._wsCall != null) {
        GoTime._wsRequestTime = Date.now();
        success = GoTime._wsCall();
        if (success) {
          GoTime._syncCount++;
          return;
        }
      }
      if (GoTime._ajaxURL != null) {
        success = GoTime._ajaxSample();
        if (success) {
          GoTime._syncCount++;
        }
      }
    };

    GoTime._calculateOffset = function(requestTime, responseTime, serverTime) {
      var duration, oneway;
      duration = responseTime - requestTime;
      oneway = duration / 2;
      return {
        offset: serverTime - requestTime + oneway,
        precision: oneway
      };
    };

    GoTime._reviseOffset = function(sample, method) {
      var now;
      if (isNaN(sample.offset) || isNaN(sample.precision)) {
        return;
      }
      now = GoTime.now();
      GoTime._lastSyncTime = now;
      GoTime._lastSyncMethod = method;
      GoTime._history.push({
        Sample: sample,
        Method: method,
        Time: now
      });
      if (sample.precision <= GoTime._precision) {
        GoTime._offset = Math.round(sample.offset);
        GoTime._precision = sample.precision;
      }
      if (!GoTime._firstSyncCallbackRan && (GoTime._firstSyncCallback != null)) {
        GoTime._firstSyncCallbackRan = true;
        return GoTime._firstSyncCallback(now, method, sample.offset, sample.precision);
      } else if (GoTime._onSyncCallback != null) {
        return GoTime._onSyncCallback(now, method, sample.offset, sample.precision);
      }
    };

    GoTime._dateFromService = function(text) {
      return new Date(parseInt(text));
    };

    return GoTime;

  })();

}).call(this);
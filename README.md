@UTF-8

# A very simple library to send events to your google analytics for GO

Analytics v0.1
=============================

The library creates an object used to send simple events to a google analytics property

First start:

Enter in your go environment and install the needed libraries:

```
go get github.com/webability-go/analytics
```

In your code, creates the object and then use it:

```
package main

import "github.com/webability-go/analytics"

var UA *analytics.UA

function main() {
  
  UA = analytics.CreateUA("[Your Universal Analytics ID]")
  
  // do something...

  // Creates a UUID if we need one
  UUID := analytics.UUID()
  UA.SetUUID(UUID)
  
  // sets geo (optional, recommended), generally based on the IP of the visitor
  UA.SetGeoID("JP")   // japan
  
  // When you need to fire an event:
  UA.Event("event-category", "event-action", "event-label", <event-value>)
  
  // do something... 
}
```

The user UUID is the unique session identifier. It should be always the same UUID for the same user. The function analytics.UUID() builds one if you need one.
If your code attends more than one user (web server or so), be sure to keep the UUID in a cookie or database session table, or so.
If your code does not attend users or session, but processes, you may use a UUID per process or even one single hardcoded UUID. 
Google uses the UUID to count unique sessions/users in analytics.

Notes on GEOID, added 2019-05-08:
-----------------------
The geoid is the ISO country 2 letters or city (see google manuals). It is set once and used when sendint the events data to google.

Notes on Multithreading:
-----------------------
If you use the analytics on a multithread server (web server for instance), every hit implies a DIFFERENT set of basic data (UUID, GeoID, etc)
so it is highly recommended to create the UA object in EACH THREAD instead of a one global, since every thread will send a differend UUID and GeoID.




TO DO:
======
- Support more parameters to send to analytics


Version Changes Control
=======================

V0.2.0 - 2019-05-08
-----------------------
- Added support for geoid parameter
- UUID set once only as geoid parameter
- Event funcion does not have anymore the UUID parameter
- Added basic test unit

V0.0.1 - 2019-04-05
-----------------------
- Creates a session and Fires a single analytics event to google

---

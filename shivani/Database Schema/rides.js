// create rideshare collection
// mongo rideshare
//insert customer 1
db.ridesDB.insert(
{
    "customerID" : "01",
    "rideID"   : "01",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "MountainView",
    "rideType"   : "VTA Train"
}
);
db.ridesDB.insert(
{
    "customerID" : "01",
    "rideID"   : "02",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "MountainView",
    "rideType"   : "Local Bus"
}
);

db.ridesDB.insert(
{
    "customerID" : "02",
    "rideID"   : "01",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "MountainView",
    "rideType"   : "VTA Train"
}
);

db.ridesDB.insert(
{
    "customerID" : "03",
    "rideID"   : "01",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "MountainView",
    "rideType"   : "Express Bus"
}
);

db.ridesDB.insert(
{
    "customerID" : "03",
    "rideID"   : "02",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "Tasman",
    "rideType"   : "VTA Train"
}
);

db.ridesDB.insert(
{
    "customerID" : "04",
    "rideID"   : "01",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "SantaClara",
    "rideType"   : "VTA Train"
}
);

db.ridesDB.insert(
{
    "customerID" : "05",
    "rideID"   : "01",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "MountainView",
    "rideType"   : "VTA Train"
}
);

db.ridesDB.insert(
{
    "customerID" : "05",
    "rideID"   : "02",
    "startTime" : ISODate("2018-03-17T04:00:00Z"),
    "startLocation" : "Paseo de SanAntonia",
    "rideType"   : "Local Bus"
}
);

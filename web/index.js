console.log("Starting app");

var platform = new H.service.Platform({
  'apikey': 'CsOaO0yHpuu-ZCgBZyvqRI-bb0F1ydmYnj_u_KbWTzg'
});
var defaultLayers = platform.createDefaultLayers();

//Instantiate (and display) a map object:
var map = new H.Map(
  document.getElementById('mapContainer'),
  defaultLayers.vector.normal.map,
  {
    zoom: 9,  
    center: { lat: 49.0477699820096, lng: 2.075042724609375 }
  }
);

async function getCoordinates() {
  let res = await axios.get("http://localhost:8080/getPath")
  let coordinates = [];
  
  for (let i = 0 ; i < res.data.features.length ; i++) {
    let feature = res.data.features[i];
    for (let j = 0 ; j < feature.geometry.coordinates.length ; j++) {
      coordinates.push(feature.geometry.coordinates[j]);
    }
  }
  
  return coordinates;
}

let coordinates = await getCoordinates();

let i = 1;
coordinates.forEach(c => {
  let svgMarkup = '<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg"><circle cx="50" cy="50" r="10"/><text x="50%" y="50%" text-anchor="middle" stroke="#51c5cf" stroke-width="2px" dy=".3em">'+ i +'</text></svg>';
  let icon = new H.map.Icon(svgMarkup);
 
  let markerCoords = {lat: c[1], lng: c[0]};
  let marker = new H.map.Marker(markerCoords, {icon: icon});

  map.addObject(marker)
  map.setCenter(markerCoords);
  i++;
});



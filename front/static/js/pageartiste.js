 // Description: Script pour la page de l'artiste
 
 
 // Variable pour stocker la carte et pour ajouter des marqueurs plus tard
 let map;

 // Chargement du DOM
 document.addEventListener("DOMContentLoaded", function () {
   // Initialiser la carte vide
   map = L.map('map').setView([0, 0], 2);
   L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
     maxZoom: 19,
     attribution: '© OpenStreetMap contributors'
   }).addTo(map);

   // Mettre a jour les données de l'artiste
   fetchArtistData();

   // Ajout des marqueurs sur la carte
   fetchLocationData();
 });

 // Récupère et met à jour les infos de l'artiste

 function fetchArtistData() {
  const url = new URL(window.location.href);
  const id = url.searchParams.get('id') || "";

  fetch(`${window.location.pathname}?id=${id}&type=artist`, {
    headers: { 'X-Requested-With': 'XMLHttpRequest' }
  })
    .then(resp => resp.json())
    .then(artist => {
      document.getElementById("artist-name").innerHTML = `<h2>${artist.Name || "Nom inconnu"}</h2>`;

      document.getElementById("artist-image").innerHTML = artist.Image
      ? `<img src="${artist.Image}" alt="Image de ${artist.Name}" class="photo-artiste" />`
      : "<p>Image non disponible</p>";    

      document.getElementById("artist-creation-date").innerHTML =
        `
        <p><strong>Date de création</strong> 
        <br>
        <br> ${artist.CreationDate || "Inconnue"}</p>`;

      document.getElementById("artist-members").innerHTML =
        `
        <p><strong>Membres</strong> 
        <br>
        <br>
        ${artist.Members?.join("<br>") || "Aucun membre renseigné"}</p>` ;

      document.getElementById("artist-first-album").innerHTML =
        `<p><strong>Premier album :</strong> ${artist.FirstAlbum || "Inconnu"}</p>`;

      // Format datesLocations
      if (artist.DatesLocations) {
        document.getElementById("artist-dates-locations").innerHTML = 
        formatDatesLocations(artist.DatesLocations);
      }
    });
}



 // Mise en forme des dates et lieux de concert
 function formatDatesLocations(datesLocations) {
   if (!datesLocations || typeof datesLocations !== "object") {
     return "<p>(Pas de dates disponibles)</p>";
   }
   let formatted = "<ul>";
   for (const [location, dates] of Object.entries(datesLocations)) {
     formatted += `
     <li><strong>${location}:</strong> ${dates.join(", ")}</li>`;
   }
   formatted += "</ul>";
   return formatted;
 }

 // Récupérer les coordonnées et les ajouter à la carte déjà initialisée
 function fetchLocationData() {
  const url = new URL(window.location.href);
  const id = url.searchParams.get('id') || "";

  fetch(`${window.location.pathname}?id=${id}&type=location`, {
    headers: { 'X-Requested-With': 'XMLHttpRequest' }
  })
    .then(response => response.json())
    .then(data => {
      // Ajouter des marqueurs sur la carte déjà existante
      const markers = data.map(coord => {
        const lat = parseFloat(coord.lat);
        const lon = parseFloat(coord.lon);
        const name = coord.display_name || "Nom inconnu";

        const marker = L.marker([lat, lon]).addTo(map);
        marker.bindPopup(`<strong>${name}</strong>`);
        return marker;
      });

      // Ajuster le zoom à tous les marqueurs
      const group = new L.featureGroup(markers);
      map.fitBounds(group.getBounds());
    });
}

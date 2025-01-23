// Description: Script pour la page home


// Récupére les éléments
const form = document.getElementById('search-form');
const inputField = document.getElementById('search-input');
const slider = document.getElementById('membersRange');
const minFieldCreationDate = document.getElementById('minCreationDate');
const maxFieldCreationDate = document.getElementById('maxCreationDate');
const minFieldFirstAlbum = document.getElementById('minFirstAlbum');
const maxFieldFirstAlbum = document.getElementById('maxFirstAlbum');
const menuSelect = document.getElementById('menu');

let timeoutId;

// Fonction pour envoyer les données au backend
function sendDataToBackend() {
  const formData = new FormData(form);
  formData.append('members', slider.value);
  formData.append('minCreationDate', minFieldCreationDate.value);
  formData.append('maxCreationDate', maxFieldCreationDate.value);
  formData.append('minFirstAlbum', minFieldFirstAlbum.value);
  formData.append('maxFirstAlbum', maxFieldFirstAlbum.value);
  formData.append('menu', menuSelect ? menuSelect.value : 'desactiver');

  fetch('/home', {
    method: 'POST',
    body: formData,
  }).then(() => {
    // Attendre avant d'appeler fetchData
    setTimeout(fetchData, 500);
  });
}


// Gestion de la recherche avec temporisation
function handleInputWithDelay() {
  clearTimeout(timeoutId);
  timeoutId = setTimeout(sendDataToBackend, 500);
}

// Gestion de l'événement d'entrée sur le champ de recherche
inputField.addEventListener('input', handleInputWithDelay);

// Gestion des événements sur les champs min et max
[minFieldCreationDate, maxFieldCreationDate, minFieldFirstAlbum, maxFieldFirstAlbum].forEach((field) => {
  field.addEventListener('input', handleInputWithDelay);
});

// Gestion de l'événement submit du formulaire
form.addEventListener('submit', function (event) {
  event.preventDefault();
  sendDataToBackend();
});

// Mise à jour de l'affichage pour le filtre des membres et envoie au backend
function updateMemberOutput(value) {
  const output = document.getElementById('membersOutput');
  output.textContent = value === "0" ? "Désactivé" : value;

    // Envoie la valeur au backend
    sendDataToBackend();
}


// Envoie de la valeur du menu au backend
menuSelect.addEventListener('change', function () {
    sendDataToBackend();
});


// Défilement des photos
function scrollback() {
  const carousel = document.querySelector('.photo-slider');
  carousel.scrollBy({ left: -200, behavior: 'smooth' });
}

function scrollnext() {
  const carousel = document.querySelector('.photo-slider');
  carousel.scrollBy({ left: 200, behavior: 'smooth' });
}

// Récupération des données (résultats dynamiques)
function fetchData() {
  fetch('/home', {
    method: 'GET',
    headers: { 'X-Requested-With': 'XMLHttpRequest' },
  })
    .then((response) => response.json())
    .then((data) => {
      const container = document.getElementById('resultat-recherche');
      let htmlContent = '';

      if (data.artists && Array.isArray(data.artists)) {
        data.artists.forEach((artist) => {
          htmlContent += `
          <a href="/pageartiste?id=${artist.ID}" class="artist">
              <div class="artist-photo">
                ${
                  artist.Image
                    ? `<img src="${artist.Image}" alt="${artist.Name}" width="100">`
                    : '<p>(Aucune image)</p>'
                }
              </div>
              <div class="artist-nom">
                <p>${artist.Name || 'N/A'}</p>
              </div>
          </a>`;
        
          
        });
      } else {
        htmlContent = '<p>Aucun artiste trouvé.</p>';
      }
      container.innerHTML = htmlContent;
    })
    .catch((error) => {
      console.error('Erreur lors de la récupération des données :', error);
      document.getElementById('resultat-recherche').innerText = 'Erreur de chargement des données.';
    });
}

// Actualisation des résultats
fetchData();

// Fonction pour le bouton home
function submitHome() {
  // Met la valeur du champ caché "home"
  document.getElementById('home-field').value = 'home';

  // Réinitialise "mot" à un string vide
  document.getElementById('search-input').value = '';

  // Envoie le formulaire
  document.getElementById('search-form').submit();
}

// Fonction pour gérer les filtres
function applyFilter(filterType) {
    // Ajouter la logique de filtrage ici
    console.log('Applying filter:', filterType);
  }

  // Fonction pour afficher les détails des photos
  function showPhotoDetails(photoId) {
    // Ajouter la logique d'affichage des détails ici
    console.log('Showing details for photo:', photoId);
  }

  // Gérer la soumission du formulaire de recherche
  document.querySelector('form').addEventListener('submit', function(e) {
    const searchInput = document.getElementById('recherche');
    if (!searchInput.value.trim()) {
      e.preventDefault();
      alert('Veuillez entrer un terme de recherche');
    }
  });

let position = 0;
const slider = document.getElementById('slider');
const photoWidth = 200; // Largeur photo + gap

function slide(direction) {
  const maxPosition = slider.children.length - 5;
  position = Math.max(0, Math.min(position + direction, maxPosition));
  slider.style.transform = `translateX(-${position * photoWidth}px)`;
}
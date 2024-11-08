// 3D Cube visualization
function displayCube(cube) {
    const cubeDisplay = document.getElementById('cube');
    cubeDisplay.innerHTML = '';

    const cube3D = document.createElement('div');
    cube3D.className = 'cube-3d';

    const numbers = document.createElement('div');
    numbers.className = 'numbers-container';
    numbers.style.transform = `rotateX(${rotationX}deg) rotateY(${rotationY}deg)`;

    // Add outer borders
    const borders = document.createElement('div');
    borders.className = 'cube-borders';
    
    // Create all 6 faces
    const faces = ['front', 'back', 'right', 'left', 'top', 'bottom'];
    faces.forEach(face => {
        const border = document.createElement('div');
        border.className = `border-face border-${face}`;
        borders.appendChild(border);
    });

    numbers.appendChild(borders);

    // Add numbers (existing code)
    const spacing = 60;
    for (let z = 0; z < 5; z++) {
        for (let y = 0; y < 5; y++) {
            for (let x = 0; x < 5; x++) {
                const number = document.createElement('div');
                number.className = 'number';
                number.textContent = cube[z][y][x];
                
                const xPos = (x - 2) * spacing;
                const yPos = (y - 2) * spacing;
                const zPos = (z - 2) * spacing;
                
                const position = `translate3d(${xPos}px, ${yPos}px, ${zPos}px)`;
                number.style.transform = `${position} rotateX(${-rotationX}deg) rotateY(${-rotationY}deg)`;
                
                numbers.appendChild(number);
            }
        }
    }

    cube3D.appendChild(numbers);
    cubeDisplay.appendChild(cube3D);
    
    // Rotation controls
    let isDragging = false;
    let previousX = 0;
    let previousY = 0;

    cube3D.addEventListener('mousedown', (e) => {
        isDragging = true;
        previousX = e.clientX;
        previousY = e.clientY;
    });

    document.addEventListener('mousemove', (e) => {
        if (isDragging) {
            const deltaX = e.clientX - previousX;
            const deltaY = e.clientY - previousY;
            
            rotationY += deltaX * 0.5;
            rotationX -= deltaY * 0.5;

            numbers.style.transform = `rotateX(${rotationX}deg) rotateY(${rotationY}deg)`;
            
            // Update all numbers to face camera
            document.querySelectorAll('.number').forEach(num => {
                const transform = num.style.transform.split('rotateX')[0];
                num.style.transform = `${transform} rotateX(${-rotationX}deg) rotateY(${-rotationY}deg)`;
            });
            
            previousX = e.clientX;
            previousY = e.clientY;
        }
    });

    document.addEventListener('mouseup', () => {
        isDragging = false;
    });
}

// Player controls
function togglePlay() {
    isPlaying = !isPlaying;
    document.getElementById('playPause').textContent = isPlaying ? 'Pause' : 'Play';
    
    if (isPlaying) {
        playbackInterval = setInterval(() => {
            if (currentStateIndex >= data.states.length - 1) {
                isPlaying = false;
                document.getElementById('playPause').textContent = 'Play';
                clearInterval(playbackInterval);
                return;
            }
            currentStateIndex++;
            displayState(currentStateIndex);
        }, 1000 / parseInt(document.getElementById('speed').value));
    } else {
        clearInterval(playbackInterval);
    }
}

// Event listeners
document.addEventListener('DOMContentLoaded', function() {
    // Hide temperature element
    const tempElement = document.getElementById('temperature');
    if (tempElement && tempElement.parentElement) {
        tempElement.parentElement.style.display = 'none';
    }
    
    // Run selector
    document.getElementById('runSelect').addEventListener('change', function(e) {
        const runNumber = e.target.value.replace('run', '');
        loadData(runNumber);
    });

    // Play/Pause button
    document.getElementById('playPause').addEventListener('click', togglePlay);

    // Progress bar
    document.getElementById('progress').addEventListener('input', function(e) {
        currentStateIndex = parseInt(e.target.value);
        displayState(currentStateIndex);
        if (isPlaying) {
            clearInterval(playbackInterval);
            togglePlay();
        }
    });

    // Speed selector
    document.getElementById('speed').addEventListener('change', function() {
        if (isPlaying) {
            clearInterval(playbackInterval);
            togglePlay();
        }
    });

    // Initial load
    loadData(1);
});
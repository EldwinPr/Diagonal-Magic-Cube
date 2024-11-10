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

    // Event listener for mouse movement
    document.addEventListener('mousemove', (e) => {
        if (isDragging) {
            // Calculate the change in mouse position
            const deltaX = e.clientX - previousX;
            const deltaY = e.clientY - previousY;
            
            // Update rotation based on mouse movement
            rotationY += deltaX * 0.5; // Adjust rotationY
            rotationX -= deltaY * 0.5; // Adjust rotationX

            // Apply rotation to 'numbers' element
            numbers.style.transform = `rotateX(${rotationX}deg) rotateY(${rotationY}deg)`;
            
            // Update all '.number' elements to face the camera
            document.querySelectorAll('.number').forEach(num => {
                // Extract current transform without rotation
                const transform = num.style.transform.split('rotateX')[0];
                // Adjust rotation to negate container's rotation, so numbers face camera
                num.style.transform = `${transform} rotateX(${-rotationX}deg) rotateY(${-rotationY}deg)`;
            });
            
            // Update previous mouse positions
            previousX = e.clientX;
            previousY = e.clientY;
        }
    });

    // Event listener for mouse up to stop dragging
    document.addEventListener('mouseup', () => {
        isDragging = false;
    });
}

// Player controls
// Function to toggle playback state
function togglePlay() {
    isPlaying = !isPlaying; // Toggle play/pause state
    // Update play/pause button text
    document.getElementById('playPause').textContent = isPlaying ? 'Pause' : 'Play';
    
    if (isPlaying) {
        // Start playback interval for animation
        playbackInterval = setInterval(() => {
            // If reached the end of data states, stop playback
            if (currentStateIndex >= data.states.length - 1) {
                isPlaying = false;
                clearInterval(playbackInterval); // Stop the interval
                return;
            }

            // Move to the next state
            currentStateIndex++;
            // Update visualization with new state
            updateVisualization(data.states[currentStateIndex]);
        }, playbackSpeed); // playbackSpeed is the interval in milliseconds
    } else {
        // If paused, clear the interval
        clearInterval(playbackInterval);
    }
}
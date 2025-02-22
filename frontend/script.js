const canvas = document.getElementById("boidsCanvas");
const ctx = canvas.getContext("2d");

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

async function fetchBoids() {
    try {
        const response = await fetch("/boids");
        const flock = await response.json();
        drawBoids(flock);
    } catch (error) {
        console.error("Failed to fetch boids:", error);
    }
}

function drawBoids(flock) {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    flock.boids.forEach(boid => {
        ctx.beginPath();
        ctx.arc(boid.x, boid.y, 5, 0, Math.PI * 2);
        ctx.fillStyle = "white";
        ctx.fill();
        ctx.closePath();
    });
}

setInterval(fetchBoids, 100);

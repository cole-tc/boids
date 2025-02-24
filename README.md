## boids
visualization of bird flocking behaviors (i.e. starlings)

*Note:* This is a WIP repo, there are a ton of improvements still to be made 
w.r.t. calculation efficiency, visual appeal, and customization.  

### Built Using
- **Backend:** Golang
- **Frontend:** HTML / JS

### Build + Run
*Prerequisites:* Golang installed

- Navigate to the `backend/` directory
- Build the executable by running `go build boids.go main.go`
  - The `-o` flag can optionally be provided to specify the output directory for the executable
- Run the executable with `./boids`
  - Or `{OUTPUT_DIR}/boids` if using `-o` flag

The app attaches to `localhost:8080`, which you can open in your web browser to see the visualization

### References + Inspiration
- **Craig Reynolds**
  - [Background](https://www.red3d.com/cwr/boids/)
  - [Flocks, Herds, and Schools:
    A Distributed Behavioral Model](https://team.inria.fr/imagine/files/2014/10/flocks-hers-and-schools.pdf)
- **Conrad Parker**
  - [Psuedocode](https://vergenet.net/~conrad/boids/pseudocode.html)
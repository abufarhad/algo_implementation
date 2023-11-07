//Mark and Sweep Algorithm

// --------------Steps----------------
// Push into HEAP A, B, C, D
// Create a reference from A to B
// Create a reference from A to C
// Create a reference from B to D
// Remove C reference from A
// Remove B reference from A
// Add B reference to D

//Run the gc(mark and sweep) algorithm


let HEAP =[]

const A ={
    language:"Javascript"
};
HEAP.push(A);


const B ={
    language:"Rust"
};
HEAP.push(B);
//Object B is reachable from A  and it is allocated in the memory
A.B = B;



const C ={
    language:"Ruby"
};
HEAP.push(C);
//Object C is reachable from A  and it is allocated in the memory
A.C = C;


// Let’s remove “C” reference from “A” and add another object “D”.
delete A.C;


const D ={
    language:"Golang"
};
HEAP.push(D);
//Object B is reachable from A  and it is allocated in the memory
B.D = D 

// "B" reference is removed from "A".
delete A.B;
// It means that "D" still has the reference to it from "B" but it's
// not reachable (because B is not reachable anymore)


const root = () => HEAP[0];

// Traverse all the reachable objects starting from the root and set the
// __mark__ bit on it to 1
const mark = () => {
    // Initially only the root is reachable
    let reachables = [ root() ];
    
    while (reachables.length) {
      let current = reachables.pop();
      // Mark the object if it is not already marked
      if (!current.__mark__) {
        current.__mark__ = 1;
        // add all the reachable objects from the current object
        // reachables array
        for (let i in current) {
            if (current !== null && typeof current[i] === 'object') {
            // Add it to the reachables
            reachables.push(current[i]);
          }
        }
      }
    }
  }


// Traverse the heap and move all unmarked or unreachable objects to the free list.
const sweep = () => {
    HEAP = HEAP.filter((current) => {
     if(current.__mark__ === 1) {
         current.__mark__ = 0 // For future Garbage collection cycles, reset the __mark__ bit to 0
         return true
     }else {
         return false // move it to the free list
     }
    });
 };


// Garbage collector (uses mark and sweep algorithm )
const gc = () => {
    // Set __mark__ bits on the reachable objects to 1
    mark();

    // Collect the garbage (objects with __mark__ bit not set to 1)
    sweep();
}

const main = () => {
    console.log("\nHeap state before garbage collection: ", HEAP);
    gc()
    console.log("\nHeap state after garbage collection: ", HEAP);
}

main();
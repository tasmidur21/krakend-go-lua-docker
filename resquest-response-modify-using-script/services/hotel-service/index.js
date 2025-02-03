const express = require("express");
const app = express();

app.use(express.json());

// Sample API Endpoint
app.get("/api/data", (req, res) => {
    console.log("header",req.headers);
    
    res.json({ message: "Hello from Backend" });
});

const PORT = 3000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));

import cors from "cors";
import express from "express";

const app = express();

app.use(cors());
app.use(express.json());

app.get("/", async(_req, res) => {
  res.json("Hello");
});

app.listen(3000, () => {
  console.log("Server is listening on port 3000");
});
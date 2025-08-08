import express from 'express';
import path from 'path';

const app = express();
const port = 3034;


app.set('views', path.join(__dirname, '../views'));
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
    let users = [
        {name: "Lothar", age: 58, premium: true},
        {name: "Janina", age: 52, premium: true},
        {name: "Susi", age: 82, premium: false},
    ]
    res.render('home', {users});
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
})
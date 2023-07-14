import React, { useEffect, useState } from 'react';
import './App.css';

function displayData(e) {
  
}

function App() {
  const [name, setName] = useState("")
  const [data, setData] = useState([])

  const handleSubmit = (event) => {
      event.preventDefault()
      
      const url = "localhost:1323/" + name
    
      const fetchInfo = () => {
        return fetch(url)
          .then((res) => res.json())
          .then((d) => setData(d))
          .catch(e => {
            console.log(e)
          })
      }
    
      useEffect(() => {
        fetchInfo();
      }, []);

      return (
        <div>
          <p>{JSON.stringify(data)}</p>
        </div>
      )
    }

    return (
    <div className="App">
      <form onSubmit={handleSubmit}>
        <label>
          Enter article name:
          <input
            type="text"
            name="daniel"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </label>
        <input type="submit"/>
      </form>
    </div>)
}

export default App;

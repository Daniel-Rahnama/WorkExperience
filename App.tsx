import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [name, setName] = useState("")
  const [data, setData] = useState([])
  const [elementName, setElementName] = useState("")
  const [elementValue, setElementValue] = useState("")

  const url = "http://127.0.0.1:1323" + window.location.pathname

  const fetchInfo = () => {
    return fetch(url)
      .then((res) => res.json())
      .then((d) => setData(d))
  }

  useEffect(() => {
    document.title = "Article DB"
    fetchInfo();
  }, [])

  if (window.location.pathname == '/') {
    
    const getArticle = (event) => {
      event.preventDefault()
      window.location.pathname = name
    }

    return (
      <form onSubmit={getArticle}>
        <label>
          Enter Article Name:
          <input 
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </label>
      </form>
    )
  }

  var d = JSON.stringify(data)
  var _d = JSON.parse(d)

  const keys: string[] = []

  for (var n in (_d?.properties?.elements)) {
    keys.push(n)
  }

  const table = keys.map((key) => {
    return <tr><th>{key}</th><th>{_d?.properties?.elements[key].value}</th></tr>
  })

  const sendElement = (event) => {
    const element = {"properties":{"name":_d?.properties?.name,"elements":{[elementName]:{"value":elementValue}}}}
    fetch(url, {
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      body: JSON.stringify(element)
    })
  }
  
  return (
    <div>
      <h1>{_d?.properties?.name}</h1>
      <table>{table}</table>
      <form onSubmit={sendElement}>
        <label>
          Submit new element:<br/>
          <input
            type="text"
            value={elementName}
            onChange={(e) => setElementName(e.target.value)}
          />
          <input
            type="text"
            value={elementValue}
            onChange={(e) => setElementValue(e.target.value)}
          />
          <button type="submit">Submit</button>
        </label>
      </form>
      <button onClick={() => {
        fetch(url, { method: 'DELETE'})
        window.location.reload()
      }}>Delete Article</button>
    </div>
  )
}

export default App;
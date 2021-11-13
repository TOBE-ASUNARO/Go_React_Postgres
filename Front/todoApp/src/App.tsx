import React, { useEffect, useState } from 'react';
import logo from './logo.svg'
import './App.css'
import axios from "axios"

function App () {
  const [count, setCount] = useState(0)
  const [usrSei, setUsrSei] = useState("");
  const [usrMei, setUsrMei] = useState("");

  useEffect(() => {
    // マウント後に実行したい処理

    axios
      .get(
        //  atode
        //  API変更にてコメントアウト
        // `http://localhost:3000/api/v1/user/${this.userData.id}/reservations/`,key_headers)
        `http://localhost:8000/`
      )
      .then((response) => {
        console.log(response.data)
        console.log(response.data[0].NameSei)
        setUsrSei(response.data[0].NameSei)
        setUsrMei(response.data[0].NameMei)
        // console.log(response.data[0])
        // console.log(response.data[0].user)
        // // Vuex store
        // this.$store.dispatch("userReservationData/update", response.data)
        // this.$store.dispatch("userReservationData/updateErr", "")
      })

      .catch((error) => {
        // TODO: 適切な Error 表示
        console.log(error.response)
        // console.log(error.response.data.error)
        // this.$store.dispatch(
        //   "userReservationData/updateErr",
        //   error.response.data.error
        // )
      })
      .finally(() => {
        console.log("axios finished")
      })
  }, []);
  
  
  
  
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello Vite + React!</p>
        <p>
          <button type="button" onClick={() => setCount((count) => count + 1)}>
            count is: {count}
          </button>
        </p>
        <p>
          Edit <code>App.tsx</code> and save to test HMR updates.
        </p>
        <p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          {' | '}
          <a
            className="App-link"
            href="https://vitejs.dev/guide/features.html"
            target="_blank"
            rel="noopener noreferrer"
          >
            Vite Docs
          </a>
        </p>
        
        <p>
          {usrSei}
        </p>

        <p>
          {usrMei}
        </p>
                
      </header>
    </div>
  )
}

export default App

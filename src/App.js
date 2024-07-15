import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer';
import UrlShortener from './components/UrlShortener';
import UrlList from './components/UrlList';

function App() {
  return (
    <div className="App">
      <Header />
      <main>
        <UrlShortener />
        <UrlList />
      </main>
      <Footer />
    </div>
  );
}

export default App;
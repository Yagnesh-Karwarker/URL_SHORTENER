import React, { useState, useEffect } from 'react';

function UrlList() {
  const [urls, setUrls] = useState([]);

  useEffect(() => {
    // Fetch the list of URLs from your backend
    // This is a placeholder and needs to be implemented in your backend
    const fetchUrls = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/urls');
        const data = await response.json();
        setUrls(data);
      } catch (error) {
        console.error('Error fetching URLs:', error);
      }
    };

    fetchUrls();
  }, []);

  return (
    <div className="url-list">
      <h2>Shortened URLs</h2>
      <ul>
        {urls.map((url) => (
          <li key={url.id}>
            <a href={`http://localhost:8080/api/${url.short_url}`} target="_blank" rel="noopener noreferrer">
              {url.short_url}
            </a>
            {' -> '}
            {url.original_url}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default UrlList;
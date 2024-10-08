async function conect () {
    const response = await fetch("http://localhost:5002/users", {
        method: 'GET',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      const responseData = await response.json();
      console.log(responseData)
}

conect();
// https://docs.cypress.io/api/table-of-contents

describe('My First Test', () => {
  it('Visits the app root url', () => {
    const host = Cypress.env('host')
    cy.visit(host + '/')
    cy.contains('h1', 'Welcome to Your Vue.js App')
  })
})

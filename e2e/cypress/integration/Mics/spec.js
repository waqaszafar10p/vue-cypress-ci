/// <reference types="cypress"/>
it('loads login page', () => {
  cy.visit('/')
  cy.contains('Welcome to Your Vue.js App').should('be.visible')
})

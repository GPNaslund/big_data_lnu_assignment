// COPIED FROM MY OWN IMPLEMENTATION IN ASSIGNMENT WT1

// Defines the scope types for dependencies.
export const SCOPES = {
  SINGLETON: 'singleton',
  TRANSIENT: 'transient'
};

// A simple Inversion of Control (IoC) container for managing dependencies.
class IoCContainer {
  constructor() {
    // Maps dependency names to their singleton instances.
    this.dependencyMap = {};

    // Stores the constructor and scope for each registered dependency.
    this.dependencyConfig = {};
  }

  // Registers a dependency with the IoC container.
  register(name, constructor, scope) {
    this.dependencyConfig[name] = { constructor, scope };
  }

  // Resolves and returns an instance of the requested dependency.
  resolve(name) {
    const config = this.dependencyConfig[name];
    if (!config) {
      throw new Error(`No dependency found for: ${name}`);
    }
    if (config.scope === SCOPES.SINGLETON) {
      if (!this.dependencyMap[name]) {
        this.dependencyMap[name] = new config.constructor();
      }
      return this.dependencyMap[name];
    } else if (config.scope === SCOPES.TRANSIENT) {
      return new config.constructor();
    } else {
      throw new Error(`Invalid scope specified for: ${name}`);
    }
  }
}

export default IoCContainer;

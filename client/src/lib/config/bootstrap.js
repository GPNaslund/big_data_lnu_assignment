import IoCContainer, { SCOPES } from '$lib/config/ioc';
import DashboardService from '$lib/service/DashboardService';
import SearchService from '$lib/service/SearchService';

// Creates a new IoCContainer
const container = new IoCContainer();

// Registers services
container.register("DashboardService", DashboardService, SCOPES.SINGLETON);
container.register("SearchService", SearchService, SCOPES.SINGLETON);

export default container;

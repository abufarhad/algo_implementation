Load balancing is a technique used in distributed systems to distribute network traffic across multiple servers. This method aims to optimize resource use, maximize throughput, minimize response time, and prevent overload of any single resource.

- Some of the most common Load balancing algorithms:
1. Round Robin: In Round Robin, the load balancer cycles through a list of servers and sends individual requests to each in turn. When it reaches the end of the list, it starts over at the beginning.
2. Weighted Round Robin: Suppose server 1 has twice as much processing power as server 2. So, for every request server 2 gets, server 1 will handle twice as many requests.
3. 
4. Least Connections: This method directs traffic to the server with the fewest active connections. 
5. Least Response Time: The load balancer sends requests to the server that has the fewest active connections and the lowest average response time.
6. 
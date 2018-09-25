/*
 * Copyright 2017-2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.api.v1.http;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.ArrayList;

@Path(HttpHealth.BASE_URI)
public class HttpHealth {
    public static final String BASE_URI = "/sys/info";

    @GET
    @Produces({MediaType.APPLICATION_JSON})
    @Path("ping")
    public Response ping() {
        return Response.status(200).build();
    }

    @GET
    @Produces({MediaType.APPLICATION_JSON})
    @Path("version")
    public Response getVersion() {
        String name = this.getClass().getPackage().getImplementationTitle();
        String version = this.getClass().getPackage().getImplementationVersion();
        return Response
            .status(200)
            .entity(new VersionResponse(name, version))
            .build();
    }

    public class VersionResponse {
        public String name;
        public String version;

        public VersionResponse(String name, String version) {
            this.name = name;
            this.version = version;
        }
    }

    @GET
    @Produces({MediaType.APPLICATION_JSON})
    @Path("health")
    public Response getHealthStatus() {
        return Response
            .status(200)
            .entity(new HealthStatus("ok", "TODO: this is a more detailed, bespoke health check", new ArrayList()))
            .build();
    }

    public class HealthStatus {
        public String status;
        public String summary;
        public ArrayList details;

        public HealthStatus(String status, String summary, ArrayList details) {
            this.status = status;
            this.summary = summary;
            this.details = details;
        }
    }

    @GET
    @Produces({MediaType.TEXT_HTML})
    @Path("metrics")
    public Response getMetrics() {
        String name = this.getClass().getPackage().getImplementationTitle();
        String version = this.getClass().getPackage().getImplementationVersion();
        String status = "ok";
        String summary = "TODO: this is a more detailed, bespoke health check";
        return Response
            .ok(
                "# api_server\n"
                    + "version{name=\""+name+"\",version=\""+version+"\"} 0\n"
                    + "health{status=\""+status+"\",summary=\""+summary+"\"} 0\n"
            )
            .build();
    }
}

/*
 * Copyright 2016-2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.controller;

import io.netty.handler.codec.http.HttpResponseStatus;
import io.vertx.core.AbstractVerticle;
import io.vertx.core.Future;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerResponse;
import io.vertx.core.json.Json;
import io.vertx.core.json.JsonObject;
import io.vertx.core.json.JsonArray;
import io.vertx.ext.web.Router;
import io.vertx.ext.web.RoutingContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HTTPServer extends AbstractVerticle {
    private static final Logger log = LoggerFactory.getLogger(HTTPServer.class);

    private HttpServer server;
    private final int port;

    public HTTPServer(int port) {
        this.port = port;
    }

    @Override
    public void start(Future<Void> startPromise) {

        Router router = Router.router(vertx);
         router.route("/").handler(request -> {
            HttpServerResponse response = request.response();
            response
                .setStatusCode(HttpResponseStatus.OK.code())
                .end();
        });
        router.get("/healthz").handler(request -> request.response().setStatusCode(HttpResponseStatus.OK.code()).end());
        router.get("/sys/info/ping").handler(this::ping);
        router.get("/sys/info/version").handler(this::getVersion);
        router.get("/sys/info/health").handler(this::getHealth);
        router.get("/sys/info/metrics").handler(this::getMetrics);

        server = vertx.createHttpServer();
        server.requestHandler(router::accept);
        server.listen(port, result -> {
            if (result.succeeded()) {
                log.info("Started HTTP server listening on {}", port);
                startPromise.complete();
            } else {
                log.info("Failed starting HTTP server: {}", result.cause().getMessage());
                startPromise.fail(result.cause());
            }
        });
    }

    @Override
    public void stop(Future<Void> stopPromise) {
        if (server != null) {
            server.close(stopPromise);
        } else {
            stopPromise.complete();
        }
    }

    private void ping(RoutingContext rc) {
        JsonObject json = new JsonObject();
       json.put("status", "ok");
       json.put("summary", "Address space controller is up and running");
        rc.response()
           .setStatusCode(200)
           .putHeader("content-type", "application/json")
           .end(Json.encodePrettily(json));
   }

    private void getVersion(RoutingContext rc) {
        JsonObject json = new JsonObject();
       json.put("name", this.getClass().getPackage().getImplementationTitle());
       json.put("version", this.getClass().getPackage().getImplementationVersion());
        rc.response()
           .setStatusCode(200)
           .putHeader("content-type", "application/json")
           .end(Json.encodePrettily(json));
   }

    private void getHealth(RoutingContext rc) {
       JsonObject json = new JsonObject();
       json.put("status", "ok");
       json.put("summary", "TODO: this is a more detailed, bespoke health check");
       json.put("details", new JsonArray());
        rc.response()
           .setStatusCode(200)
           .putHeader("content-type", "application/json")
           .end(Json.encodePrettily(json));
   }

    private void getMetrics(RoutingContext rc) {
        String name = this.getClass().getPackage().getImplementationTitle();
        String version = this.getClass().getPackage().getImplementationVersion();
        rc.response()
            .setStatusCode(200)
            .putHeader("content-type", "text/html")
            .end(
                "# Address space controller\n"
                    + "version{name=\""+name+"\",version=\""+version+"\"} 0\n"
                    +  "health{status=\"ok\",summary=\"TODO: this is a more detailed, bespoke health check\"} 0\n"
            );
    }
}

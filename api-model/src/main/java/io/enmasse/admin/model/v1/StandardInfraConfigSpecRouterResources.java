/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.admin.model.v1;

import java.util.Objects;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;

import io.fabric8.kubernetes.api.model.Doneable;
import io.sundr.builder.annotations.Buildable;
import io.sundr.builder.annotations.Inline;

@Buildable(
        editableEnabled = false,
        generateBuilderPackage = false,
        builderPackage = "io.fabric8.kubernetes.api.builder",
        inline = @Inline(type = Doneable.class, prefix = "Doneable", value = "done")
)
@JsonPropertyOrder({"cpu", "memory"})
@JsonInclude(JsonInclude.Include.NON_NULL)
public class StandardInfraConfigSpecRouterResources extends AbstractWithAdditionalProperties {
    private String memory;
    private String cpu;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        StandardInfraConfigSpecRouterResources that = (StandardInfraConfigSpecRouterResources) o;
        return Objects.equals(memory, that.memory) &&
                Objects.equals(cpu, that.cpu);
    }

    @Override
    public int hashCode() {
        return Objects.hash(memory, cpu);
    }

    @Override
    public String toString() {
        return "StandardInfraConfigSpecRouterResources{" +
                "memory='" + memory + '\'' +
                ", cpu='" + cpu + '\'' +
                '}';
    }

    public void setMemory(String memory) {
        this.memory = memory;
    }

    public String getMemory() {
        return memory;
    }

    public String getCpu() {
        return cpu;
    }

    public void setCpu(String cpu) {
        this.cpu = cpu;
    }

}

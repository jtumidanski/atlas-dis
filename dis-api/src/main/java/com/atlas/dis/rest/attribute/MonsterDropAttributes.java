package com.atlas.dis.rest.attribute;

import rest.AttributeResult;

public record MonsterDropAttributes(Integer monsterId, Integer itemId, Integer maximumQuantity, Integer minimumQuantity,
                                    Integer chance) implements AttributeResult {
}
